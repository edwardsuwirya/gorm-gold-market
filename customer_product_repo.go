package main

import "fmt"

type CustomerProductRepo struct {
	BaseRepo
}

func (cr *CustomerProductRepo) Insert(newCustomerProduct CustomerProduct) error {
	result := cr.conn.Db.Create(&newCustomerProduct)
	return cr.HandleError(result)
}

func (cr *CustomerProductRepo) FindTotalCustomerByProduct() AggProductResult {
	var total AggProductResult

	subQuery1 := cr.conn.Db.Model(&CustomerProducts{}).Select("customer_product_id", "COUNT(customer_product_id) as total ").Group("customer_product_id")
	result := cr.conn.Db.Table("(?) as CustomerProductCount", subQuery1).
		Select("mst_customer_product.first_name as product_name,CustomerProductCount.total").
		Joins("JOIN mst_customer_product ON mst_customer_product.id = CustomerProductCount.customer_product_id").
		Scan(&total)
	//	result := cr.conn.Db.Raw(`
	//		SELECT mst_customer_product.first_name as product_name,CustomerProductCount.total
	//        FROM
	//		(SELECT customer_product_id,COUNT(customer_product_id) as total
	//		FROM customer_products
	//		GROUP BY customer_product_id) CustomerProductCount JOIN mst_customer_product ON
	//			mst_customer_product.id = CustomerProductCount.customer_product_id
	//`).Scan(&total)
	fmt.Println(total)
	err := cr.HandleError(result)
	if err != nil {
		//	return AggResult{
		//		Total: -1,
		//	}
		return AggProductResult{
			ProductName: "",
			Total:       -1,
		}
	}
	return total
}

func NewCustomerProductRepo(conn *DbConn) *CustomerProductRepo {
	return &CustomerProductRepo{
		BaseRepo{
			conn: conn,
		},
	}
}
