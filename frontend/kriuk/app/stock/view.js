"use client"
import { DataTable } from "@/components/ui/data-table"
import { KriukPopup } from "@/components/ui/kriuk-popup"
import { Popup } from "@/components/ui/popup"
import { useState } from "react"

const StockView = ({data}) => {
    const [stocks, setStocks] = useState(data)


    const handleSave = () => {

    }

    

    const colDefs = [
        {
            accessorKey: "name",
            header: "Name",
        },
        {
            accessorKey: "qty",
            header: "Quantity",
        },
        {
            accessorKey: "unit",
            header: "Unit",
        },

        {
            accessorKey: "price",
            header: "Price",
        },
        {
            accessorKey: "createon",
            header: "Create On",
        },
        {
            accessorKey: "updateon",
            header: "Update On",
        },
        {
            id: "actions",
            header: "Actions",
            cell: ({ row }) => {
                return (
                  <div>
                    <EditDialog
                      title={"Edit Stock"}
                      btnLabel={"Edit"}
                      data={row.original}
                    />
                  </div>
                );
            }
        }
    ]

    return (
        <div>
            <h1 className="text-2xl font-bold mb-4">Stocks</h1>
            <DataTable columns={colDefs} data={stocks}/>
        </div>
    )
}

const EditDialog = ({title, btnLabel, data}) => {
    const [stockData, setStockData] = useState({
      id: data.id,
      name: data.name,
      qty: data.qty,
      unit: data.unit,
      price: data.price,
    });

    const handleEdit = (data) => {
      console.log(data);
    };

    return (
      <KriukPopup btnTxt={btnLabel} headerTxt={title} onSave={()=> handleEdit(stockData)}>
        <div className="flex gap-2">
          <div className="grid gap-4 py-y">
            <div className="grid grid-cols-4 items-center gap-4">
              <label htmlFor="name" className="text-right text-sm">
                Name
              </label>
              <input
                id={stockData.name}
                value={stockData["name"]}
                onChange={(e) =>
                  setStockData({ ...stockData, ["name"]: e.target.value })
                }
                className="col-span-3 p-2 border rounded bg-gray-900 text-white"
              />
            </div>
            <div className="grid grid-cols-4 items-center gap-4">
              <label htmlFor="quantity" className="text-right text-sm">
                Quantity
              </label>
              <input
                id={stockData.qty}
                value={stockData["qty"]}
                onChange={(e) =>
                  setStockData({ ...stockData, ["qty"]: e.target.value })
                }
                className="col-span-3 p-2 border rounded bg-gray-900 text-white"
              />
            </div>
          </div>
        </div>
      </KriukPopup>
    );
}

export default StockView