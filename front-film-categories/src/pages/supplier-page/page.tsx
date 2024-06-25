import { useState } from "react";
import { useSuppliers } from "../../hooks/suppliers";
import { Button } from "../../components/ui/button";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "../../components/ui/table";
import { AddSupplierForm } from "./components/add-supplier-form";
import { UpdateSupplierForm } from "./components/update-supplier-form";
import PopUpDialog from "../../components/pop-up-dialog";

type OpenDialogs = {
  [key: string]: boolean;
};

export default function SuppliersPage() {
  const { suppliers, deleteSupplier, getSuppliers } = useSuppliers();
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);
  const [openDialogs, setOpenDialogs] = useState<OpenDialogs>({});

  const setIsUpdateDialogOpen = (id: string, isOpen: boolean) => {
    setOpenDialogs((prev) => ({ ...prev, [id]: isOpen }));
  };


  const handleCreateDialogClose = () => {
    setIsCreateDialogOpen(false);
  };

  return (
    <div className="space-y-8 flex flex-col">
      <div className="p-8 flex items-center">
        <h1 className="text-2xl font-semibold border-b mr-4">Categories</h1>
        <PopUpDialog
          isOpen={isCreateDialogOpen}
          onOpenChange={setIsCreateDialogOpen}
          title="Add category"
          text="Add"
          FormComponent={
            <AddSupplierForm
              onClose={handleCreateDialogClose}
              onRefresh={getSuppliers}
            />
          }
        />
      </div>
      <div className="flex items-center justify-center">
        <div className="w-full max-w-screen-sm mx-auto">
          <Table className="w-full">
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Email</TableHead>
                <TableHead>Phone</TableHead>
                <TableHead className="text-left">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {suppliers ? (
                suppliers.map((supplier) => (
                  <TableRow key={supplier.id}>
                    <TableCell>{supplier.name}</TableCell>
                    <TableCell>{supplier.email}</TableCell>
                    <TableCell>{supplier.phone}</TableCell>
                    <TableCell className="text-left">
                      <PopUpDialog
                        isOpen={openDialogs[supplier.id] || false}
                        onOpenChange={(isOpen) => setIsUpdateDialogOpen(supplier.id, isOpen)}
                        title="Edit category"
                        text="Edit"
                        FormComponent={
                          <UpdateSupplierForm
                            onClose={() => setIsUpdateDialogOpen(supplier.id, false)}
                            onRefresh={getSuppliers}
                            id={supplier.id}
                            name={supplier.name}
                            email={supplier.email}
                            phone={supplier.phone}
                          />
                        }
                      />
                      <Button
                        variant="destructive"
                        onClick={() => deleteSupplier(supplier.id)}
                        className="ml-2"
                      >
                        Delete
                      </Button>
                    </TableCell>
                  </TableRow>
                ))
              ) : (
                <TableRow>
                  <TableCell colSpan={2}>Loading...</TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </div>
      </div>
    </div>
  );
}
