// import { useState } from "react";
// import FilterSection from "./FilterSection";
import "./index.scss";
// import SearchInput from "../widgets/SearchInput";
import { Typography } from "@mui/material";
import { SearchBar } from "../SearchBar";

/**
 * TODO: Advanced structured search
 *
 * Design:
 * - Single input
 * - Show "Advanced Search" on focus
 * - Support structured filters:
 *   $tag:xxx
 *   $category:xxx
 * - Selected filters appear as removable tokens inside input
 *
 * Similar to GitHub / Notion search
 */

const ArticleFilterPanel = () => {
  // const [category, setCategory] = useState("all");
  // const [tag, setTag] = useState("all");
  // const [order, setOrder] = useState("default");

  return (
    <div className="article-filter-panel">
      <Typography className="article-list-title">Article Search</Typography>

      {/* <div className="search-input"> */}
      <SearchBar></SearchBar>
      {/* </div> */}

      {/* 
      <FilterSection
        title="分类"
        value={category}
        onChange={setCategory}
        options={[
          { label: "全部", value: "all" },
          { label: "golang", value: "golang" },
          { label: "博客教程", value: "blog" },
          { label: "网站配置", value: "site" },
        ]}
      />

      <FilterSection
        title="标签"
        value={tag}
        onChange={setTag}
        options={[
          { label: "全部", value: "all" },
          { label: "golang", value: "golang" },
          { label: "关于作者", value: "author" },
          { label: "广告合作", value: "ads" },
        ]}
      />

      <FilterSection
        title="排序"
        value={order}
        onChange={setOrder}
        options={[
          { label: "默认", value: "default" },
          { label: "时间", value: "time" },
          { label: "评论", value: "comment" },
          { label: "浏览", value: "view" },
          { label: "点赞", value: "like" },
        ]}
      />
       */}
    </div>
  );
};

export default ArticleFilterPanel;
