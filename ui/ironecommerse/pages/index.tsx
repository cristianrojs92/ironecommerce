import Head from 'next/head'
import Main from '../src/layout/main/main' 
import * as service_products from '../src/services/products'
import * as service_gategories from '../src/services/categories'


export default function Home(props) {

  return (
    <div className="main-content">
      <Head>

        <title>Iron-Ecommerse</title>
       
        <link rel="icon" href="/favicon.ico" />
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css"/>
        <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet"/>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>

      </Head>

      <main className = "cristian">
        <Main products={props.products} categories={props.categories} products_offers={props.products_offers}/>
      </main>

    </div>
  )
}

export async function getStaticProps() {

  //Obtenemos los productos
  const products = await service_products.getAll();

  //Obtenemos las ofertas
  const products_offers = await service_products.getOffers();

  //Obtenemos las categorias
  const categories = await service_gategories.getAll();
  
  return {
    props: {
      products,
      products_offers,
      categories
   }    
  }

}