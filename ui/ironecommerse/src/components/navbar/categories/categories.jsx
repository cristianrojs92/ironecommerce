/*
 * categories.jsx
 *
 * Created on 5 de Junio de 2020
 * Copyright (c) 2020
 * Author Cristian Rojas <b>cristiarojs92@gmail.com</b>
 *
 */

import navbarStyles from '../../../../styles/components/navbar/navbar.module.css';
import {useEffect } from 'react';

export default function Categories(props) {

  useEffect(() => {

    //Manejamos la lista de categorias
    var options_categories = {inDuration:300,outDuration:225,constrainWidth:!1,hover:false,gutter:0,coverTrigger:false,alignment:"right",closeOnClick : false}
    handleShowLiist('.dropdown-trigger', options_categories);


    let options = {accordion:true}
      var elems = document.querySelectorAll('.collapsible');
      var instances = M.Collapsible.init(elems, options);
   

      
  }); 

  //Esta funcion se encarga de listar elementos al seleccionar un item
  function handleShowLiist(name_class,options) { 

    var elems = document.querySelectorAll(name_class);
    
    M.Dropdown.init(elems, options);
  }



  //Esta funcion se encarga de filtrar las categorias pricipales para que sean visualizadas en el menu
  const showMainCategories = (() => {

    //Obtenemos las categorias de las props
    const categories = props.categories;

    if(categories !== undefined && categories.length > 0) {

      //Filtramos las categorias principales
      const main_categories = categories.filter(cat => cat.main == true)

      if(main_categories !== undefined && main_categories.length > 0) {
         return (
           <ul  className="collapsible">
            { main_categories !== undefined ? main_categories.map( c =>
            <li key={c.id}>
              <div className="collapsible-header">
                <i className="material-icons">{c.icon_name}</i>
                <span className="category-title">{c.name}</span>
                
              </div>
              
              <div className="collapsible-body">
                <ul className="collection">
                   
                  {  (c.subcategories !== null) ?  c.subcategories.map(sub =>  <li key={sub.id} className="collection-item">{sub.name}</li> ) : ''}
                </ul>
              </div>
            </li>) : ''}
          </ul>)
      }
    }
  });

  return (
      <li>
          <a className='dropdown-trigger' href='#' data-target='categories' >Categorias</a>
          <div className={`${navbarStyles["categories-container"]}`}>
            <ul id='categories' className={`${navbarStyles["list-categories-container"]} dropdown-content list-container`}>
              { 
                showMainCategories()
              }
            </ul>
          </div>
      </li>
  )
}
