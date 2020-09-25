/*
 * categories.d.ts
 *
 * Created on 6 de Junio de 2020
 * Copyright (c) 2020
 * Author Cristian Rojas <b>cristiarojs92@gmail.com</b>
 *
 */

declare namespace ironecommerse.categories {

  export interface category {
    id: number;
    main: bolean;
    name: string;
    icon_name : string;
    subcategories : category[];
  }
  
}