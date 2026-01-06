import { useState } from "react";
import { login, type LoginRequest } from "@/api/user";
import { useSnackbar } from "notistack";
import { useUserStore } from "@/stores/user";
import { auth } from "@/utils/auth";

export const useLogin = () => {
  const [loading, setLoading] = useState(false);
  const { enqueueSnackbar } = useSnackbar();
  const setUser = useUserStore((s) => s.setUser);

  const submit = async (data: LoginRequest) => {
    setLoading(true);

    try {
      const res = await login(data);
      if (res.code === 0) {
        enqueueSnackbar("登录成功", { variant: "success" });
        setUser(res.data.user);
        auth.setAuth(res.data.access_token, res.data.access_token_expires_at);
        return true;
      } else {
        enqueueSnackbar(res.msg || "登录失败", { variant: "error" });
        return false;
      }
    } catch {
      enqueueSnackbar("网络错误", { variant: "error" });
      return false;
    } finally {
      setLoading(false);
    }
  };

  return { submit, loading };
};
