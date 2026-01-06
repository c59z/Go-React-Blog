import { Avatar, IconButton, Popover } from "@mui/material";
import { useState } from "react";
import type { MouseEvent } from "react";
import { useUserStore } from "@/stores/user";
import "./index.scss";
import { useNavigate } from "react-router-dom";
import { auth } from "@/utils/auth";
import { logout } from "@/api/user";

const UserAvatar = () => {
  const userInfo = useUserStore((s) => s.user);
  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>(null);
  const navigate = useNavigate();

  const open = Boolean(anchorEl);

  const handleOpen = (e: MouseEvent<HTMLElement>) => {
    setAnchorEl(e.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleLogin = () => {
    navigate("/login");
  };

  const handleLogout = async () => {
    try {
      await logout();
    } finally {
      auth.logout();
      navigate("/", { replace: true });
    }
  };

  const loggedBox = (
    <div className="logged-panel">
      <div className="logged-header">
        <Avatar
          className="avatar"
          src={`${import.meta.env.VITE_SERVER_URL}${
            userInfo?.avatar || undefined
          }`}
        >
          {userInfo?.username?.[0]?.toUpperCase()}
        </Avatar>
        <div className="info">
          <div className="name">{userInfo?.username}</div>
          <div className="email">{userInfo?.email}</div>
        </div>
      </div>

      <div className="logged-actions">
        <button className="action-item">
          <span className="icon">👤</span>
          <span>Profile</span>
        </button>

        <button className="action-item">
          <span className="icon">🛠</span>
          <span>Management</span>
        </button>
      </div>

      <div className="logged-footer">
        <button onClick={handleLogout} className="logout-btn">
          Logout
        </button>
      </div>
    </div>
  );

  const guestBox = (
    <div className="guest-panel">
      <div className="guest-actions">
        <button onClick={handleLogin} className="action-item">
          <span className="icon">👤</span>
          <span className="text">登录</span>
        </button>

        <button className="action-item">
          <span className="icon">✏️</span>
          <span className="text">注册</span>
        </button>

        <button className="action-item">
          <span className="icon">🔒</span>
          <span className="text">找回密码</span>
        </button>
      </div>

      <div className="guest-divider">快速登录</div>

      <div className="guest-oauth">
        <button className="oauth-btn github">
          <span className="icon">🐙</span>
          <span>GitHub 登录</span>
        </button>
      </div>
    </div>
  );

  return (
    <>
      <IconButton className="user-avatar-button" onClick={handleOpen}>
        <Avatar
          className="avatar"
          src={`${import.meta.env.VITE_SERVER_URL}${
            userInfo?.avatar || undefined
          }`}
        >
          {userInfo?.username?.[0]?.toUpperCase()}
        </Avatar>
      </IconButton>

      <Popover
        open={open}
        anchorEl={anchorEl}
        disableScrollLock
        onClose={handleClose}
        className="user-avatar-popover"
        anchorOrigin={{ vertical: "bottom", horizontal: "right" }}
        transformOrigin={{ vertical: "top", horizontal: "right" }}
      >
        <div className="user-panel">{userInfo ? loggedBox : guestBox}</div>
      </Popover>
    </>
  );
};

export default UserAvatar;
