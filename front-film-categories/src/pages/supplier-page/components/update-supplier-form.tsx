import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"
import { useForm } from "react-hook-form"
import { useSuppliers } from "../../../hooks/suppliers";
import { BaseFormProps } from "../../../types"
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

interface UpdateSupplierFormProps extends BaseFormProps {
  id: string;
  name: string;
  email: string;
  phone: string;
}

const formSchema = z.object({
  name: z.string().min(1, {
    message: "Name is required",
  }).max(150, {
    message: "Name is too long",
  }),
  email: z.string().email({
    message: "Invalid email",
  }),
  phone: z.string().min(1, {
    message: "Phone is required",
  }).max(15, {
    message: "Invalid phone number",
  }),
})

export function UpdateSupplierForm({ onClose, onRefresh, id, name, email, phone }: UpdateSupplierFormProps) {
  const { updateSupplier } = useSuppliers();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: name,
      email: email,
      phone: phone,
    },
  })

  async function onSubmit(values: z.infer<typeof formSchema>) {
    await updateSupplier(id, values);
    onRefresh();
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
                Name of the supplier.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input placeholder="" {...field} />
              </FormControl>
              <FormDescription>
                Email of the supplier.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="phone"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Phone</FormLabel>
              <FormControl>
                <Input placeholder="" {...field} />
              </FormControl>
              <FormDescription>
                Phone of supplier.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <div className="flex justify-between">
          <Button type="button" variant="ghost" onClick={onClose}>Close</Button>
          <Button type="submit">Create</Button>
        </div>
      </form>
    </Form>
  )
}