package internal

import (
	"database/sql"

	models "github.com/ivanArevalo1997/HackthonGo/internal/models"
)

type RepositoryCustomer interface {
	//GetOne(id int) (models.Customer, error)
	StoreAllCustomers() ([]models.Customer, error)
}

type repositoryCustomer struct {
	db *sql.DB
}

func NewRepositoryCustomers(database *sql.DB) RepositoryCustomer {
	return &repositoryCustomer{db: database}
}

// func (repo *repositoryProducto) GetOne(id int) (models.Customer, error) {
// 	repoCat := internal.NewRepositoryCategoria(repo.db)
// 	// query := "select p.id, p.nombre, p.precio, c.id, c.nombre from producto p inner join categoria c on p.idcategoria = c.id WHERE p.id = ?"
// 	query := "SELECT id, nombre, precio, idcategoria FROM Producto WHERE id = ?"

// 	var ProductoLeido models.Producto

// 	rows, err := repo.db.Query(query, id)

// 	if err != nil {
// 		return ProductoLeido, err
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&ProductoLeido.ID, &ProductoLeido.Nombre, &ProductoLeido.Precio, &ProductoLeido.Categoria.ID)
// 		// err = rows.Scan(&ProductoLeido.ID, &ProductoLeido.Nombre, &ProductoLeido.Precio, &ProductoLeido.Categoria.ID, &ProductoLeido.Categoria.Nombre)
// 		if err != nil {
// 			return ProductoLeido, err
// 		}
// 		ProductoLeido.Categoria, err = repoCat.GetOne(ProductoLeido.Categoria.ID)
// 		if err != nil {
// 			return ProductoLeido, err
// 		}
// 	}
// 	return ProductoLeido, nil
// }

func (repo *repositoryCustomer) StoreAllCustomers() ([]models.Customer, error) {
	var tablaSumasPorCategorias []models.Customer
	// query := "select p.id, p.nombre, p.precio, c.id, c.nombre from producto p inner join categoria c on p.idcategoria = c.id WHERE p.id = ?"
	query := "insert c.nombre, sum(p.precio) suma from producto p inner join categoria c on p.idcategoria = c.id group by c.nombre order by suma"

	rows, err := repo.db.Query(query)

	if err != nil {
		return tablaSumasPorCategorias, err
	}

	for rows.Next() {
		var sumaAgregar models.Customer
		// err = rows.Scan(&ProductoLeido.ID, &ProductoLeido.Nombre, &ProductoLeido.Precio, &ProductoLeido.Categoria.ID, &ProductoLeido.Categoria.Nombre)
		if err != nil {
			return tablaSumasPorCategorias, err
		}

		tablaSumasPorCategorias = append(tablaSumasPorCategorias, sumaAgregar)
	}

	return tablaSumasPorCategorias, nil
}
