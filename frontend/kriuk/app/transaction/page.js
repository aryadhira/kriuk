"use client"
import { DataTable } from "@/components/ui/data-table"

const Transaction = () => {
    const colDefs = [
            {
                accessorKey: "id",
                header: "ID",
            },
            {
                accessorKey: "name",
                header: "Name",
            },
            {
                accessorKey: "email",
                header: "Email",
            },
            {
                accessorKey: "status",
                header: "Status",
                cell: ({ row }) => (
                <span className="px-2 py-1 text-xs rounded-full bg-green-900 text-green-100">
                    {row.original.status}
                </span>
                ),
            },
        ]

    const rowData = [
        { id: 1, name: "John Doe", email: "john@example.com", status: "Active" },
        { id: 2, name: "Jane Smith", email: "jane@example.com", status: "Inactive" },
        { id: 3, name: "Bob Johnson", email: "bob@example.com", status: "Active" },
    ]
    return (
        <div>
            <h1 className="text-2xl font-bold mb-4">Transactions</h1>
            <DataTable columns={colDefs} data={rowData} />
        </div>
    )
}

export default Transaction;