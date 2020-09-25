/*
 * products.d.ts
 *
 * Created on 5 de Junio de 2020
 * Copyright (c) 2020
 * Author Cristian Rojas <b>cristiarojs92@gmail.com</b>
 *
 */

declare namespace ironecommerse.products {

  export interface product {
    id: number;
    id_category: number;
    name: string;
    description: string;
    price: number;
    stock: number;
    image: string;
  }
  
}