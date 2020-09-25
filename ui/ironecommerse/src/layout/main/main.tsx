import Navbar from '../../components/navbar/navbar'
import Carousel from '../../components/carousel/carousel'
import Recomended from '../../components/recommended/recommended'

export default function Main(props) {
  return (
    <div>
      <Navbar categories = {props.categories}/>
      <Carousel/>
      <Recomended products_offers = {props.products_offers}/>
    </div>
  )
}
