import { Box, Button, TextField, CircularProgress } from "@mui/material";
import { useState } from "react";
import { useLogin } from "./hook/useLogin";
import { useCaptcha } from "./hook/useCaptcha";
import "./index.scss";
import { useNavigate } from "react-router-dom";

const LoginPage = () => {
  const { submit, loading } = useLogin();
  const { captchaId, img, refresh } = useCaptcha();
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [captcha, setCaptcha] = useState("");

  const handleLogin = async () => {
    const ok = await submit({
      email,
      password,
      captcha,
      captcha_id: captchaId,
    });

    if (!ok) {
      refresh();
      setCaptcha("");
      return;
    }

    navigate("/", { replace: true });
  };

  return (
    <Box className="login-page">
      <Box className="login-card">
        <h1 className="login-title">Sign in</h1>
        <p className="login-subtitle">Welcome back</p>

        <Box className="login-form">
          <TextField
            label="邮箱"
            fullWidth
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />

          <TextField
            label="密码"
            type="password"
            fullWidth
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />

          <Box className="captcha-row">
            <TextField
              label="验证码"
              value={captcha}
              onChange={(e) => setCaptcha(e.target.value)}
            />
            <img src={img} alt="captcha" onClick={refresh} />
          </Box>

          <Button
            fullWidth
            variant="contained"
            onClick={handleLogin}
            disabled={loading}
            className="login-button"
          >
            {loading ? <CircularProgress size={22} /> : "登录"}
          </Button>

          <div className="login-footer">
            <span>还没有账号？</span>
            <a className="register-link">注册</a>
          </div>
        </Box>
      </Box>
    </Box>
  );
};

export default LoginPage;
