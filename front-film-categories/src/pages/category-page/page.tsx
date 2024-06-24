import { useState } from "react";
import { useCategories } from "../../hooks/categories";
import { Button } from "../../components/ui/button";
import { Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "../../components/ui/dialog";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "../../components/ui/table";
import { AddCategoryForm } from "./components/add-category-form";

export default function CategoriesPage() {
  const { categories, deleteCategory } = useCategories();
  const [isDialogOpen, setIsDialogOpen] = useState(false);

  const handleDialogClose = () => {
    setIsDialogOpen(false);
  };

  return (
    <div className="space-y-8 flex flex-col">
      <div className="p-8 flex items-center">
        <h1 className="p-4 text-2xl font-semibold border-b">Categories</h1>
        <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
          <DialogTrigger asChild className=" ml-4">
            <Button variant="default">Add</Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[425px]">
            <DialogHeader>
              <DialogTitle>Add category</DialogTitle>
            </DialogHeader>
            <DialogContent>
              <AddCategoryForm onClose={handleDialogClose} />
            </DialogContent>
            <DialogFooter>
              <Button type="submit" form="add-category-form">Save changes</Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </div>
      <div className="flex items-center justify-center">
        <div className="w-full max-w-screen-md mx-auto">
          <Table className="w-full">
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {categories ? (
                categories.map((category) => (
                  <TableRow key={category.id}>
                    <TableCell>{category.name}</TableCell>
                    <TableCell>
                      <Button
                        className="mr-2"
                        variant="outline"
                        onClick={() => { }}
                      >
                        Editar
                      </Button>
                      <Button
                        variant="destructive"
                        onClick={() => deleteCategory(category.id)}
                      >
                        Excluir
                      </Button>
                    </TableCell>
                  </TableRow>
                ))
              ) : (
                <TableRow>
                  <TableCell colSpan={4}>Loading...</TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </div>
      </div>
    </div>
  );
}
