// components/DataTable.jsx
"use client"

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { Button } from "@/components/ui/button"
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationPrevious,
  PaginationNext,
} from "@/components/ui/pagination"

import {
  flexRender,
  getCoreRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
} from "@tanstack/react-table"

import { useState } from "react"

export function DataTable({ columns, data }) {
  const [pageSize, setPageSize] = useState(5)
  const [pageIndex, setPageIndex] = useState(0)

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    manualPagination: true,
    pageCount: Math.ceil(data.length / pageSize),
    state: {
      pagination: { pageIndex, pageSize },
    },
    onPaginationChange: (updater) => {
      const newState = typeof updater === "function" ? updater({ pageIndex, pageSize }) : updater
      setPageIndex(newState.pageIndex)
      setPageSize(newState.pageSize)
    },
  })

  const canPreviousPage = pageIndex > 0
  const canNextPage = pageIndex < table.getPageCount() - 1

  const getPageNumbers = () => {
    const totalPages = table.getPageCount()
    const currentPage = pageIndex
    const pages = []

    if (totalPages <= 5) {
      for (let i = 0; i < totalPages; i++) pages.push(i)
    } else {
      if (currentPage <= 2) {
        pages.push(0, 1, 2, 'ellipsis', totalPages - 1)
      } else if (currentPage >= totalPages - 3) {
        pages.push(0, 'ellipsis', totalPages - 3, totalPages - 2, totalPages - 1)
      } else {
        pages.push(0, 'ellipsis', currentPage, 'ellipsis', totalPages - 1)
      }
    }
    return pages
  }

  return (
    <div className="rounded-md border border-gray-700 bg-gray-900 text-white shadow-sm overflow-hidden">
      {/* Table */}
      <div className="w-full text-xs overflow-auto max-h-96">
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow key={headerGroup.id} className="hover:bg-gray-800">
                {headerGroup.headers.map((header) => (
                  <TableHead key={header.id} className="text-white font-semibold">
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                          header.column.columnDef.header,
                          header.getContext()
                        )}
                  </TableHead>
                ))}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
                  className="hover:bg-gray-800 transition-colors"
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id}>
                      {flexRender(cell.column.columnDef.cell, cell.getContext())}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="h-24 text-center"
                >
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>

      {/* Pagination */}
      <div className="flex flex-col sm:flex-row justify-between items-center px-4 py-3 bg-gray-800 border-t border-gray-700 gap-4">
        <div className="text-xs text-gray-300">
          Page {pageIndex + 1} of {table.getPageCount()}
        </div>

        <Pagination className={'text-xs'}>
          <PaginationContent>
            <PaginationItem>
              <PaginationPrevious
                onClick={() => table.previousPage()}
                className={canPreviousPage ? "cursor-pointer" : "cursor-not-allowed opacity-50"}
              />
            </PaginationItem>

            {getPageNumbers().map((page, i) => {
              if (page === 'ellipsis') {
                return (
                  <PaginationItem key={i}>
                    <span className="text-gray-400 px-2">...</span>
                  </PaginationItem>
                )
              }
              return (
                <PaginationItem key={i}>
                  <PaginationLink
                    href="#"
                    onClick={(e) => {
                      e.preventDefault()
                      table.setPageIndex(page)
                    }}
                    isActive={pageIndex === page}
                  >
                    {page + 1}
                  </PaginationLink>
                </PaginationItem>
              )
            })}

            <PaginationItem>
              <PaginationNext
                onClick={() => table.nextPage()}
                className={canNextPage ? "cursor-pointer" : "cursor-not-allowed opacity-50"}
              />
            </PaginationItem>
          </PaginationContent>
        </Pagination>

        <div className="text-xs text-gray-300 hidden sm:block">
          {table.getPaginationRowModel().rows.length} of {data.length} rows
        </div>
      </div>
    </div>
  )
}