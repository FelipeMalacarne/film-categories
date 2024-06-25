import { useState } from "react";
import { useCategories } from "../../hooks/categories";
import { Button } from "../../components/ui/button";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "../../components/ui/table";
import { AddCategoryForm } from "./components/add-category-form";
import { UpdateCategoryForm } from "./components/update-category-form";
import PopUpDialog from "../../components/pop-up-dialog";
import { SquarePen, Trash2 } from "lucide-react";

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
        <div className="flex-1 space-y-4 p-8 pt-6">
            <div className="flex items-center justify-between space-y-2">
                <h2 className="text-3xl font-bold tracking-tight">Categories</h2>
                <PopUpDialog
                    isOpen={isCreateDialogOpen}
                    onOpenChange={setIsCreateDialogOpen}
                    title="Add category"
                    text="Add a new category"
                    FormComponent={
                        <AddCategoryForm
                            onClose={handleCreateDialogClose}
                            onRefresh={getCategories}
                        />}
                />
            </div>
            <Table className="w-full">
                <TableHeader>
                    <TableRow>
                        <TableHead>ID</TableHead>
                        <TableHead>Name</TableHead>
                        <TableHead>Created At</TableHead>
                        <TableHead className="text-left">Actions</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {categories ? (
                        categories.map((category) => (
                            <TableRow key={category.id}>
                                <TableCell>{category.id}</TableCell>
                                <TableCell>{category.name}</TableCell>
                                <TableCell>{new Date(category.created_at).toLocaleString('pt-BR')}</TableCell>
                                <TableCell className="text-left">
                                    <PopUpDialog
                                        isOpen={openDialogs[category.id] || false}
                                        onOpenChange={(isOpen) => setIsUpdateDialogOpen(category.id, isOpen)}
                                        title="Update category"
                                        text=<SquarePen/>
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
                                        className="ml-2"
                                    >
                                        <Trash2/>
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
    );
}
