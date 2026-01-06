import { captcha } from "@/api/base";
import { useEffect, useState } from "react";

export const useCaptcha = () => {
  const [captchaId, setCaptchaId] = useState("");
  const [img, setImg] = useState("");
  const [loading, setLoading] = useState(false);

  const refresh = async () => {
    setLoading(true);
    try {
      const res = await captcha();
      if (res.code === 0) {
        setCaptchaId(res.data.captcha_id);
        setImg(res.data.pic_path);
      }
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    refresh();
  }, []);

  return {
    captchaId,
    img,
    refresh,
    loading,
  };
};
