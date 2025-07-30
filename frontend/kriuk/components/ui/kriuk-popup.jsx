"use client"

import { useState } from "react"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "./dialog"
import { Button } from "./button"

export const KriukPopup = ({btnTxt, headerTxt, onSave = () => {}, children}) => {
    const [open, setOpen] = useState(false)

    const handleSubmit = () => {
        onSave()
        setOpen(false)
    }

    return (
      <Dialog open={open} onOpenChange={setOpen}>
        <DialogTrigger asChild>
          <Button variant="outline" size="sm">
            {btnTxt}
          </Button>
        </DialogTrigger>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>{headerTxt}</DialogTitle>
          </DialogHeader>
          {children}
          <div className="flex justify-end gap-2">
            <Button variant="outline" onClick={() => setOpen(false)}>
              Cancel
            </Button>
            <Button onClick={handleSubmit}>Save</Button>
          </div>
        </DialogContent>
      </Dialog>
    );
}