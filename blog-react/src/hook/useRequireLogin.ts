// hooks/useRequireLogin.ts
import { useSnackbar } from "notistack";
import { auth } from "@/utils/auth";
import { useUserStore } from "@/stores/user";

export const useRequireLogin = () => {
  const { enqueueSnackbar } = useSnackbar();
  const clearUser = useUserStore((u) => u.clearUser);

  return (action: () => void) => {
    if (!auth.isLogin()) {
      enqueueSnackbar("请先登录", { variant: "warning" });
      clearUser();
      return;
    }
    action();
  };
};
