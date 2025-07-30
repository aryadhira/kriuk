export const getStocks = async() => {
    const res = await fetch('http://localhost:9099/stocks')

    if (!res.ok){
        toast("Failed to get stock",{
            description: "get stock failed",
            action:{}
        })
        throw new Error('Failed to fetch stock')
    }

    return res.json()
}