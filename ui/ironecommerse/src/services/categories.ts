/*
 * categories.ts
 *
 * Created on 6 de Junio de 2020
 * Copyright (c) 2020
 * Author Cristian Rojas <b>cristiarojs92@gmail.com</b>
 *
 */

import * as http from '../lib/http';
import * as config from '../../config'

export async function getAll() : Promise<ironecommerse.categories.category[] | undefined> {

  let categories : ironecommerse.categories.category[] = undefined

  //Realizamos la peticion de productos a la API
  const result = await http.get(config.API_HOST, Number(config.API_PORT) ,'/categories',null);

  //Si obtuvimos un OK
  if(result.code === 200) {

    //Verificamos le resultado
    if(result.data !== undefined){

      //Inicializamos el array
      categories = [];

      for(const item of result.data) {

        const category  : ironecommerse.categories.category = {
          id: item.id,
          main: item.main,
          name: item.name,
          icon_name: item.icon_name,
          subcategories : item.subcategories
        }

        //Agregamos la categoria
        categories.push(category);
      }
      console.log(`[categories.ts] [getAll] Se obtuvieron las categorias correctamente`);
    }
  } else {
    console.log(`[categories.ts] [getAll] Error al obtener las categorias code = ${result.code}`);
  }

  return categories;

}
