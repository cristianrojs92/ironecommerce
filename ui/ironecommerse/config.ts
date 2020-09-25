/*
 * http.ts
 *
 * Created on 5 de Junio de 2020
 * Copyright (c) 2020
 * Author Cristian Rojas <b>cristiarojs92@gmail.com</b>
 *
 */


// Realiza un trace de las solicitudes enviadas y recibidas por el servicio API
export const API_TRACE = process.env.API_TRACE || true;
export const API_TRACE_BODY = process.env.API_TRACE_BODY || true;
export const API_HOST = process.env.API_HOST || 'localhost';
export const API_PORT  = process.env.API_PORT || 3001;


