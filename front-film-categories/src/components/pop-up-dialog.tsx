import React from 'react';
import { Dialog, DialogTrigger, DialogContent, DialogHeader, DialogTitle } from './ui/dialog';
import { Button } from './ui/button';

interface PopUpDialogProps {
  onOpenChange: (isOpen: boolean) => void;
  isOpen: boolean;
  title: string;
  text: string;
  FormComponent: React.ReactElement;
}

const PopUpDialog: React.FC<PopUpDialogProps> = ({
  isOpen,
  onOpenChange,
  title,
  text,
  FormComponent,
}) => {
  return (
    <Dialog open={isOpen} onOpenChange={onOpenChange}>
      <DialogTrigger asChild className="ml-4">
        <Button variant="default">{text}</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>{title}</DialogTitle>
        </DialogHeader>
        <DialogContent>
          {FormComponent}
        </DialogContent>
      </DialogContent>
    </Dialog>
  );
};

export default PopUpDialog;