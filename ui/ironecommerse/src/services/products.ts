/*
 * products.ts
 *
 * Created on 6 de Junio de 2020
 * Copyright (c) 2020
 * Author Cristian Rojas <b>cristiarojs92@gmail.com</b>
 *
 */

import * as http from '../lib/http';
import * as config from '../../config'

export async function getAll() : Promise<ironecommerse.products.product[] | undefined> {

  let products : ironecommerse.products.product[] = undefined
  //Realizamos la peticion de productos a la API
  const result = await http.get(config.API_HOST, Number(config.API_PORT) ,'/products',null);

  //Si obtuvimos un OK
  if(result.code === 200) {

    //Verificamos le resultado
    if(result.data !== undefined){

      //Inicializamos el array
      products = [];

      for(const item of result.data) {

        const product  : ironecommerse.products.product = {
          id: item.id,
          id_category: item.id_category,
          name: item.name,
          description: item.description,
          price: item.price,
          stock: item.stock,
          image: item.image
        }

        //Agregamos el producto
        products.push(product);
      }
      console.log(`[products.ts] [getAll] Se obtuvieron los productos correctamente`);
    }
  } else {
    console.log(`[products.ts] [getAll] Error al obtener los productos code = ${result.code}`);
  }

  return products;

}

export async function getOffers() : Promise<ironecommerse.products.product[] | undefined> {

  let products : ironecommerse.products.product[] = undefined
  //Realizamos la peticion de productos a la API
  const result = await http.get(config.API_HOST, Number(config.API_PORT) ,'/products/offers',null);

  //Si obtuvimos un OK
  if(result.code === 200) {

    //Verificamos le resultado
    if(result.data !== undefined){

      //Inicializamos el array
      products = [];

      for(const item of result.data) {

        const product  : ironecommerse.products.product = {
          id: item.id,
          id_category: item.id_category,
          name: item.name,
          description: item.description,
          price: item.price,
          stock: item.stock,
          image: item.image
        }

        //Agregamos el producto
        products.push(product);
      }
      console.log(`[products.ts] [getOffers] Se obtuvieron los productos correctamente`);
    }
  } else {
    console.log(`[products.ts] [getOffers] Error al obtener los productos code = ${result.code}`);
  }

  return products;

}