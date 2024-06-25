import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { Button } from "../../../components/ui/button"
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../../../components/ui/form"
import { Input } from "../../../components/ui/input"

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

export function AddFilmForm() {
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

  function onSubmit(values: z.infer<typeof formSchema>) {
    console.log(values)
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Username</FormLabel>
              <FormControl>
                <Input placeholder="" {...field} />
              </FormControl>
              <FormDescription>
                The name of the movie.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit">Submit</Button>
      </form>
    </Form>
  )
}