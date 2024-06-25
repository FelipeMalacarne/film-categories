import { useState } from "react";
import { useCategories } from "../../hooks/categories";
import { Button } from "../../components/ui/button";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "../../components/ui/table";
import { AddCategoryForm } from "./components/add-category-form";
import { UpdateCategoryForm } from "./components/update-category-form";
import PopUpDialog from "../../components/pop-up-dialog";

type OpenDialogs = {
  [key: string]: boolean;
};

export default function CategoriesPage() {
  const { categories, deleteCategory, getCategories } = useCategories();
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
        <h1 className="p-4 text-2xl font-semibold border-b">Categories</h1>
        <PopUpDialog
          isOpen={isCreateDialogOpen}
          onOpenChange={setIsCreateDialogOpen}
          title="Add category"
          text="Add"
          FormComponent={<AddCategoryForm onClose={handleCreateDialogClose} onRefresh={getCategories} />}
        />
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
                      <div className="flex space-x-2">
                        <PopUpDialog
                          isOpen={openDialogs[category.id] || false}
                          onOpenChange={(isOpen) => setIsUpdateDialogOpen(category.id, isOpen)}
                          title="Edit category"
                          text="Edit"
                          FormComponent={
                            <UpdateCategoryForm
                              onClose={() => setIsUpdateDialogOpen(category.id, false)}
                              onRefresh={getCategories}
                              id={category.id}
                              name={category.name}
                            />
                          }
                        />
                        <Button
                          variant="destructive"
                          onClick={() => deleteCategory(category.id)}
                        >
                          Delete
                        </Button>
                      </div>
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
