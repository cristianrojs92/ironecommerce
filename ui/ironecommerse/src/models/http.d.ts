/*
 * httpResponses.d.ts
 *
 * Created on 5 de Junio de 2020
 * Copyright (c) 2020
 * Author Cristian Rojas <b>cristiarojs92@gmail.com</b>
 *
 */

declare namespace ironecommerse.http {

  export interface response {
    ret: boolean;
    code?: status,
    data?: any;
    errorMsg?: string;
  }
  
}