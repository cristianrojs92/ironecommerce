/*
 * recomended.jsx
 *
 * Created on 10 de Junio de 2020
 * Copyright (c) 2020
 * Author Cristian Rojas <b>cristiarojs92@gmail.com</b>
 *
 */

import recomendedStyles from '../../../styles/components/recomended/recomended.module.css';

export default function Recomended(props) {

  return (
    <div className = "main-section">
       <div className="card-panel">
        <h5 className="center-align grey-text text-darken-4">Ofertas de la semana</h5>
        <div className="row" >
          {props.products_offers === undefined ? '' : 
            props.products_offers.map(product => 
          <div key={product.id} className="col s12 l2">
            <div className="card animate card-radius">
              <div className="card-content">
                <span className ={`${recomendedStyles[".card-title"]}`}>{(product.name).substring(0,15)}</span>
                <img src={product.image} className="responsive-img" alt=""/>
                <div>
                  <h6>${product.price}</h6>
                </div>
              </div>
            </div>
          </div>    
      )}
        </div>
      </div>
    </div>

  )
}
