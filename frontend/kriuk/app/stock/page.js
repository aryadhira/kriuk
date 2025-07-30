import { getStocks } from "@/lib/api/stock.get"
import StockView from "./view"


const Stock = async() => {
    const stocks = await getStocks()
    
    return <StockView data={stocks.data}/>
}

export default Stock;