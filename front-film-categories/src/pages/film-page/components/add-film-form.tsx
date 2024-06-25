import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"
import { useForm } from "react-hook-form"
import { useFilms } from "../../../hooks/films"
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


const formSchema = z.object({
  name: z.string().min(1, {
    message: "Name is required",
  }).max(150, {
    message: "Name is too long",
  }),
  author: z.string().min(1, {
    message: "Author is required",
  }).max(150, {
    message: "Name is too long",
  }),
  description: z.string().max(150, {
    message: "Description is too long",
  }),
  duration: z.number().int().positive(),
  release_date: z.date(),
})

export function AddFilmForm({ onClose, onRefresh }: BaseFormProps) {
  const { createFilm } = useFilms();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      author: "",
      description: "",
      duration: 0,
      release_date: new Date(),
    },
  })

  async function onSubmit(values: z.infer<typeof formSchema>) {
    await createFilm(values);
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
                Name of the film.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="author"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input placeholder="" {...field} />
              </FormControl>
              <FormDescription>
                Name of the movie author.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="description"
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
        <FormField
          control={form.control}
          name="duration"
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