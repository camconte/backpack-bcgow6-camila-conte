package products

import (
	"clase1/internal/domain"
	"clase1/pkg/utils"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	ERROR_FORZADO = errors.New("Error forzado")
)

var product_test = domain.Product{
	Id: 1,
	Name: "New Product",
	ProductType: "etc",
	Count: 11,
	Price: 120,
	WarehouseId: 1,
}

/*Si desde el repo verifican algun error real de sql, deberían utilizar ese error en especifico a retornar en el mock para llegar a verificar ese caso en particular.
Si le aplican un wrapper personalizado al error desde la capa repo, tendrían que verificar con ese ultimo error en el assert.*/

func TestStoreProduct(t *testing.T) {
	//arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Store Ok", func(t *testing.T) {

		//como se utiliza el Prepare y el Exec en el metodo, debemos mockear ambos
		mock.ExpectPrepare(regexp.QuoteMeta(STORE_PRODUCT))
		//NewResult nos devuelve un objeto que contiene el ultimo id (1) y la cantidad de rows afectadas (1) 
		mock.ExpectExec(regexp.QuoteMeta(STORE_PRODUCT)).WillReturnResult(sqlmock.NewResult(1, 1))

		columns := []string{"id", "name", "type", "count", "price"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(product_test.Id, product_test.Name, product_test.ProductType, product_test.Count, product_test.Price)

		mock.ExpectQuery(regexp.QuoteMeta(GET_BY_NAME)).WithArgs(product_test.Name).WillReturnRows(rows)

		repository := NewRepository(db)

		//act
		/* -------------------------------------------------------------------------- */
		/* LAS SENTENCIAS NO TIENEN NADA QUE VER UNA CON LA OTRA, NO SE PERSISTE 
			EN UN STORAGE, LA SENTENCIA DE STORE DEVUELVE LO QUE NOSOTROS 
			LE INDICAMOS(NEW RESULT CON EL LAST ID Y ROWS AFFECTED)                   */
		/* -------------------------------------------------------------------------- */
		

		newID, err := repository.Store(product_test)
		assert.NoError(t, err)

		productResult, err := repository.GetByName(product_test.Name)
		assert.NoError(t, err)

		//assert
		assert.NotNil(t, productResult)
		assert.Equal(t, product_test.Id, newID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Store Fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta(STORE_PRODUCT))
		//en el willReturn... nosotros determinamos qué queremos que devuelva la bd mockeada
		mock.ExpectExec(regexp.QuoteMeta(STORE_PRODUCT)).WillReturnError(ERROR_FORZADO)

		repository := NewRepository(db)

		//act
		id, err := repository.Store(product_test)

		//assert
		assert.Error(t, err)
		assert.Empty(t, id)

		//mock.ExpectationsWereMet() valida que todos los expects declarados se hayan ejecutado, incluyendo los 'WithArgs' donde espera que la funcion del mock reciba los
		// valores exactos que declaraste desde el test
		//verificar que en ninguna de todas los expect que definimos en el mock halla ocurrido algún error.
		assert.NoError(t, mock.ExpectationsWereMet())
	})
} 

func TestGetByName(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	//le indicamos el formato de las rows que va a devolver
	columns := []string{"id", "name", "type", "count", "price"}
	rows := sqlmock.NewRows(columns)

	//le indicamos la fila que queremos que devuelva la sentencia
	rows.AddRow(product_test.Id, product_test.Name, product_test.ProductType, product_test.Count, product_test.Price)
	
	mock.ExpectQuery(regexp.QuoteMeta(GET_BY_NAME)).WithArgs(product_test.Name).WillReturnRows(rows)

	repo := NewRepository(db)
	//act
	productResult, err := repo.GetByName(product_test.Name)

	//assert
	assert.NoError(t, err)
	assert.Equal(t, product_test.Name, productResult.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
} 

func TestUpdateProduct(t *testing.T) {
	//arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(UPDATE_PRODUCT))
	mock.ExpectExec(regexp.QuoteMeta(UPDATE_PRODUCT)).WithArgs(product_test.Name, product_test.ProductType, product_test.Count, product_test.Price, product_test.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewRepository(db)

	//act
	err = repository.Update(product_test, product_test.Id)
	
	
	//assert
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	
} 

func TestDeleteProduct(t *testing.T) {
	//arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(DELETE_PRODUCT))
	mock.ExpectExec(regexp.QuoteMeta(DELETE_PRODUCT)).WithArgs(product_test.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	repository := NewRepository(db)

	//act
	err = repository.Delete(product_test.Id)
	
	
	//assert
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	
} 



/* --------------------------- EJEMPLOS SQLMOCK CON MOVIE --------------------------- */

/* func TestGetOneWithContext(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
	rows := sqlmock.NewRows(columns)

	rows.AddRow(movie_test.ID, movie_test.Title, movie_test.Rating, movie_test.Awards, movie_test.Length, movie_test.Genre_id)
	mock.ExpectQuery(regexp.QuoteMeta(GET_MOVIE)).WithArgs(movie_test.ID).WillReturnRows(rows)

	repo := NewRepository(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	movieResult, err := repo.GetMovieWithContext(ctx, movie_test.ID)
	assert.NoError(t, err)
	assert.Equal(t, movie_test.Title, movieResult.Title)
	assert.Equal(t, movie_test.ID, movieResult.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
} */

/* func TestExistMovieOK(t *testing.T) {
	//creamos el mock de la base de datos
	db, mock, err := sqlmock.New()
	//chequeamos que no haya un error
	assert.NoError(t, err)
	defer db.Close()

	//declaramos las columnas que le vamos a pasar
	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	//el valor que le vamos a pasar es el que tiene que recibir el metodo exists en este caso
	rows.AddRow(movie_test.ID)

	//usamos el ExpectQuery porque en el metodo del repo se utiliza el QueryRow, si se utiliza el prepare tendriamos que utilizar ExpectPrepare
	//el regexp.QuoteMeta nos evita los signos de pregunta de la sentencia
	//el WithArgs recibe los parametros necesarios para la consulta, es el argunmento que pide en la query del exists (id=?), reemplazo de los ?
	//WillReturnRows recibe las columnas que necesitamos que devuelva
	mock.ExpectQuery(regexp.QuoteMeta(EXIST_MOVIE)).WithArgs(1).WillReturnRows(rows)

	repo := NewRepository(db)

	resp := repo.Exists(context.TODO(), 1)

	assert.True(t, resp)
} */

/* func TestExistMovieFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	mock.ExpectQuery(regexp.QuoteMeta(EXIST_MOVIE)).WithArgs(2).WillReturnRows(rows)

	repo := NewRepository(db)

	resp := repo.Exists(context.TODO(), 2)

	assert.False(t, resp)
	assert.NoError(t, mock.ExpectationsWereMet())
} */

/* func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Store Ok", func(t *testing.T) {

		//como se utiliza el Prepare y el Exec en el metodo, debemos mockear ambos
		mock.ExpectPrepare(regexp.QuoteMeta(SAVE_MOVIE))
		//NewResult nos devuelve un objeto que contiene el ultimo id (1) y la cantidad de rows afectadas (1) 
		mock.ExpectExec(regexp.QuoteMeta(SAVE_MOVIE)).WillReturnResult(sqlmock.NewResult(1, 1))

		columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(movie_test.ID, movie_test.Title, movie_test.Rating, movie_test.Awards, movie_test.Length, movie_test.Genre_id)
		mock.ExpectQuery(regexp.QuoteMeta(GET_MOVIE)).WithArgs(1).WillReturnRows(rows)

		repository := NewRepository(db)
		ctx := context.TODO()

		newID, err := repository.Save(ctx, movie_test)
		assert.NoError(t, err)

		movieResult, err := repository.GetMovieByID(ctx, int(newID))
		assert.NoError(t, err)

		assert.NotNil(t, movieResult)
		assert.Equal(t, movie_test.ID, movieResult.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Store Fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta(SAVE_MOVIE))
		//en el willReturn... nosotros determinamos qué queremos que devuelva la bd mockeada
		mock.ExpectExec(regexp.QuoteMeta(SAVE_MOVIE)).WillReturnError(ERROR_FORZADO)

		repository := NewRepository(db)
		ctx := context.TODO()

		id, err := repository.Save(ctx, movie_test)

		assert.EqualError(t, err, ERRORFORZADO.Error())
		assert.Empty(t, id)

		//mock.ExpectationsWereMet() valida que todos los expects declarados se hayan ejecutado, incluyendo los 'WithArgs' donde espera que la funcion del mock reciba los
		// valores exactos que declaraste desde el test
		//verificar que en ninguna de todas los expect que definimos en el mock halla ocurrido algún error.
		assert.NoError(t, mock.ExpectationsWereMet())
	})
} */

/* func Test_RepositoryGetAllOK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Columns
	columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
	rows := sqlmock.NewRows(columns)

	//movies seria el datamock
	movies := []domain.Movie{{ID: 1, Title: "Avatar", Rating: 22, Awards: 99, Length: 0, Genre_id: 1}, {ID: 2, Title: "Simpson", Rating: 33, Awards: 11, Length: 2, Genre_id: 2}}

	for _, movie := range movies {
		rows.AddRow(movie.ID, movie.Title, movie.Rating, movie.Awards, movie.Length, movie.Genre_id)
	}

	mock.ExpectQuery(regexp.QuoteMeta(GET_ALL_MOVIES)).WillReturnRows(rows)

	repo := NewRepository(db)
	resultMovies, err := repo.GetAll(context.TODO())

	assert.NoError(t, err)
	assert.Equal(t, movies, resultMovies)
	assert.NoError(t, mock.ExpectationsWereMet())
} */

/* func Test_RepositoryGetAllFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// Columns
	columns := []string{"id", "title", "rating", "awards", "length", "genre_id"}
	rows := sqlmock.NewRows(columns)
	movies := []domain.Movie{{ID: 1, Title: "Avatar", Rating: 22, Awards: 99, Length: 0, Genre_id: 1}, {ID: 2, Title: "Simpson", Rating: 33, Awards: 11, Length: 2, Genre_id: 2}}

	for _, movie := range movies {
		rows.AddRow(movie.ID, movie.Title, movie.Rating, movie.Awards, movie.Length, movie.Genre_id)
	}

	mock.ExpectQuery(regexp.QuoteMeta(GET_ALL_MOVIES)).WillReturnError(ERRORFORZADO)

	repo := NewRepository(db)
	resultMovies, err := repo.GetAll(context.TODO())

	assert.EqualError(t, err, ERRORFORZADO.Error())
	assert.Empty(t, resultMovies)
	assert.NoError(t, mock.ExpectationsWereMet())
} */


/* ---------------------------------- TXDB (no aplicable con Fury) --------------------------------- */
func TestStoreProductTXDB_Ok(t *testing.T) {
	//arrange
	db := utils.InitTxDB()
	defer db.Close()

	repo := NewRepository(db)

	productExpected := domain.Product{
		Name: "Samsung Smart TV 65' 4K",
		ProductType: "televisores",
		Count: 10,
		Price: 250000,
		WarehouseId: 1,
	}

	//act
	id, err := repo.Store(productExpected)
	assert.NoError(t, err)

	productExpected.Id = id
	productResult, err := repo.GetByName(productExpected.Name)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, productResult)
	assert.Equal(t, productExpected.Id, productResult.Id)
} 



/* func Test_RepositorySave_txdb(t *testing.T) {
	//seria necesario crear el package utils para probar este test
	db := utils.InitTxDB()
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.TODO()
	// (&m.Title, &m.Rating, &m.Awards, &m.Length, &m.Genre_id

	movieExp := domain.Movie{
		Title:        "Título ficticio",
		Rating:       2,
		Awards:       3,
		Length:       3,
		Genre_id:     2,
		Release_date: "2022-11-09 00:00:00",
	}

	// Act
	id, err := repo.Save(ctx, movieExp)
	assert.NoError(t, err)

	movieExp.ID = int(id)
	movies, err := repo.GetMovieByID(context.TODO(), int(id))

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, movies)
	assert.NoError(t, err)
	assert.Equal(t, movieExp.ID, movies.ID)
} */