package main

// import (
// 	"gorm.io/gorm"
// 	"gorm.io/driver/postgres"
// )

// type Customer struct{
// 	gorm.Model
//     id int64
//     name string
//     address string
// }

// type CustomerLicense struct {
// 	gorm.Model
// 	id int64
// 	customer_id int64
// 	license_id int64
// 	license_code string
// 	active bool
// 	start_date
// }
// table! {
//     customer_license (id) {
//         id -> Int4,
//         customer_id -> Int4,
//         license_id -> Int4,
//         license_code -> Varchar,
//         active -> Bool,
//         start_date -> Timestamp,
//         end_date -> Timestamp,
//         cpu -> Int4,
//         storage -> Int4,
//         number_of_nodes -> Int4,
//     }
// }

// table! {
//     license (id) {
//         id -> Int4,
//         name -> Varchar,
//         duration -> Int4,
//         cpu -> Int4,
//         storage -> Int4,
//         number_of_nodes -> Int4,
//     }
// }

// joinable!(customer_license -> customer (customer_id));
// joinable!(customer_license -> license (license_id));

// allow_tables_to_appear_in_same_query!(
//     customer,
//     customer_license,
//     license,
// );
