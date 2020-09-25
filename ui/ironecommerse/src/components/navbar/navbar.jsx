/*
 * navbar.jsx
 *
 * Created on 5 de Junio de 2020
 * Copyright (c) 2020
 * Author Cristian Rojas <b>cristiarojs92@gmail.com</b>
 *
 */

import navbarStyles from '../../../styles/components/navbar/navbar.module.css';
import Categories from './categories/categories'

export default function Navbar(props) {

  return (
    <div>
      <nav>
        <div className={`nav-wrapper ${navbarStyles["nav-background"]}`}>
          <div className={navbarStyles["nav-items-container"]}>
            <ul id="nav-mobile" className="left hide-on-med-and-down">
              <Categories categories={props.categories}/>
              <li><a >Ofertas</a></li>
            </ul>
          </div>
          <div className={navbarStyles["nav-items-container"]}>
            <ul id="nav-mobile" className="right hide-on-med-and-down">
              <li><a href="badges.html">Contactanos</a></li>
            </ul>
          </div>
        </div>
      </nav>
    </div>

  )
}
