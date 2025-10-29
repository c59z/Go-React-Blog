import { Avatar, IconButton, Popover } from "@mui/material";
import { useState } from "react";
import type { MouseEvent } from "react";
import { useUserStore } from "@/stores/user";
import "./index.scss";

const UserAvatar = () => {
  const userInfo = useUserStore((s) => s.user);
  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>(null);

  const open = Boolean(anchorEl);

  const handleOpen = (e: MouseEvent<HTMLElement>) => {
    setAnchorEl(e.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const loggedBox = (
    <div className="logged-panel">
      <div className="logged-header">
        <Avatar className="avatar" />
        <div className="info">
          <div className="name">Username</div>
          <div className="email">admin@example.com</div>
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
        <button className="logout-btn">Logout</button>
      </div>
    </div>
  );

  const guestBox = (
    <div className="guest-panel">
      <div className="guest-actions">
        <button className="action-item">
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
        <Avatar className="user-avatar-image" />
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
