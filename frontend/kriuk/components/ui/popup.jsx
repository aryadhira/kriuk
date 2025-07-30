"use client"

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { useState } from "react"

export function Popup({ rowData, onSave }) {
  const [open, setOpen] = useState(false)
  const [data, setData] = useState({ ...rowData })

  const handleSubmit = () => {
    onSave?.(data)
    setOpen(false)
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant="outline" size="sm">Edit</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Edit Item</DialogTitle>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          {Object.keys(data).map((key) => (
            <div key={key} className="grid grid-cols-4 items-center gap-4">
              <label htmlFor={key} className="text-right text-sm">
                {key}
              </label>
              <input
                id={key}
                value={data[key]}
                onChange={(e) => setData({ ...data, [key]: e.target.value })}
                className="col-span-3 p-2 border rounded bg-gray-900 text-white"
              />
            </div>
          ))}
        </div>
        <div className="flex justify-end gap-2">
          <Button variant="outline" onClick={() => setOpen(false)}>
            Cancel
          </Button>
          <Button onClick={handleSubmit}>Save</Button>
        </div>
      </DialogContent>
    </Dialog>
  )
}