import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { useFilms } from "../../../hooks/films";
import { useCategories } from "../../../hooks/categories";
import { BaseFormProps } from "../../../types";
import { Button } from "../../../components/ui/button";
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "../../../components/ui/form";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "../../../components/ui/select";

interface UpdateFilmCategoryFormProps extends BaseFormProps {
    id: string;
    categoryId: string;
}

const formSchema = z.object({
    categoryId: z.string().nonempty(),
});

export function UpdateFilmCategoryForm({ onClose, onRefresh, id, categoryId }: UpdateFilmCategoryFormProps) {
    const { categories } = useCategories();
    const { updateFilmCategory } = useFilms();

    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            categoryId: categoryId,
        },
    });

    async function onSubmit(values: z.infer<typeof formSchema>) {
        await updateFilmCategory(id, values.categoryId);
        onRefresh();
        if (onClose) {
            onClose();
        }
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                <FormField
                    control={form.control}
                    name="categoryId"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Category</FormLabel>
                            <FormControl>
                                <Select
                                    onValueChange={field.onChange}
                                    value={field.value}
                                >
                                    <SelectTrigger className="w-full">
                                        <SelectValue placeholder="Select a category" />
                                    </SelectTrigger>
                                    <SelectContent>
                                        {categories && categories.map((category) => (
                                            <SelectItem key={category.id} value={category.id}>
                                                {category.name}
                                            </SelectItem>
                                        ))}
                                    </SelectContent>
                                </Select>
                            </FormControl>
                            <FormDescription>
                                Select the category for the film.
                            </FormDescription>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <div className="flex justify-between">
                    <Button type="button" variant="ghost" onClick={onClose}>Close</Button>
                    <Button type="submit">Update</Button>
                </div>
            </form>
        </Form>
    );
}
