import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"
import { useForm } from "react-hook-form"
import { useCategories } from "../../../hooks/categories"
import { Button } from "../../../components/ui/button"
import { Input } from "../../../components/ui/input"
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../../../components/ui/form"

interface AddCategoryFormProps {
  onClose: () => void;
}

const formSchema = z.object({
  name: z.string().min(1, {
    message: "Name is required",
  }).max(150, {
    message: "Name is too long",
  }),
})

export function AddCategoryForm({ onClose }: AddCategoryFormProps) {
  const { createCategory } = useCategories();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
    },
  })

  async function onSubmit(values: z.infer<typeof formSchema>) {
    await createCategory(values);
    if (onClose) { onClose() }
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input placeholder="" {...field} />
              </FormControl>
              <FormDescription>
                Category name.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit">Create</Button>
      </form>
    </Form>
  )
}