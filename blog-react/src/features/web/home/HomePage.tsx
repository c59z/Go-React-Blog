import { captcha } from "@/api/base";
import { useEffect } from "react";

const HomePage = () => {
  useEffect(() => {
    captcha()
      .then((res) => {
        console.log(res.data);
      })
      .catch(() => {});
  }, []);

  return <>HomePage</>;
};

export default HomePage;
