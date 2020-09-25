/*
 * carousel.jsx
 *
 * Created on 10 de Junio de 2020
 * Copyright (c) 2020
 * Author Cristian Rojas <b>cristiarojs92@gmail.com</b>
 *
 */

import {useEffect } from 'react';

export default function Carousel(props) {

  useEffect(() => {

    var elems = document.querySelectorAll('.carousel');
    let options = {duration:200,fullWidth:true,indicators:true}
    var instances = M.Carousel.init(elems, options);
  
  }); 


  return (
    <div className="carousel carousel-slider">
      <a className="carousel-item" href="#one!"><img src='/slice/1.webp'/></a>
      <a className="carousel-item" href="#two!"><img src='/slice/2.webp'/></a>
      <a className="carousel-item" href="#three!"><img src='/slice/3.webp'/></a>
  </div>
  )
}
