import { ref } from "vue";
import { defineStore } from "pinia";

export type ToastVariant = "success" | "error" | "info";

export interface ToastItem {
  id: number;
  message: string;
  variant: ToastVariant;
}

let nextToastId = 1;

export const useToastStore = defineStore("toast", () => {
  const toasts = ref<ToastItem[]>([]);

  function remove(id: number) {
    toasts.value = toasts.value.filter((toast) => toast.id !== id);
  }

  function show(message: string, variant: ToastVariant = "info") {
    const id = nextToastId++;
    toasts.value = [...toasts.value, { id, message, variant }];

    window.setTimeout(() => {
      remove(id);
    }, 4000);
  }

  function success(message: string) {
    show(message, "success");
  }

  function error(message: string) {
    show(message, "error");
  }

  function info(message: string) {
    show(message, "info");
  }

  return {
    toasts,
    remove,
    show,
    success,
    error,
    info,
  };
});
