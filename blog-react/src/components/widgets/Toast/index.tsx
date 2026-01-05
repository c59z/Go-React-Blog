import { useCallback, useState } from "react";
import { Snackbar, Alert } from "@mui/material";

export type ToastType = "success" | "error" | "info" | "warning";

interface ToastState {
  open: boolean;
  message: string;
  type: ToastType;
}

export const useToast = () => {
  const [toastState, setToastState] = useState<ToastState>({
    open: false,
    message: "",
    type: "success",
  });

  const show = useCallback((message: string, type: ToastType) => {
    setToastState({
      open: true,
      message,
      type,
    });
  }, []);

  const hide = useCallback(() => {
    setToastState((prev) => ({
      ...prev,
      open: false,
    }));
  }, []);

  const Toast = (
    <Snackbar
      open={toastState.open}
      autoHideDuration={3000}
      onClose={hide}
      anchorOrigin={{ vertical: "top", horizontal: "center" }}
    >
      <Alert
        onClose={hide}
        severity={toastState.type}
        variant="filled"
        sx={{ width: "100%" }}
      >
        {toastState.message}
      </Alert>
    </Snackbar>
  );

  return {
    toast: {
      success: (msg: string) => show(msg, "success"),
      error: (msg: string) => show(msg, "error"),
      info: (msg: string) => show(msg, "info"),
      warning: (msg: string) => show(msg, "warning"),
    },
    Toast,
  };
};
