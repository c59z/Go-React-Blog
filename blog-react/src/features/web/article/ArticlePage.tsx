import TwoColumnLayout from "@/components/layout/TwoColumnLayout";
import "./ArticlePage.scss";
import ArticleContent from "./components/ArticleContent";
import ArticleTOC from "./components/ArticleTOC";
import { useState } from "react";

export interface HeadingItem {
  id: string;
  text: string;
  level: number;
}

const ArticleDetailPage = () => {
  const [headings, setHeadings] = useState<HeadingItem[]>([]);

  return (
    <div className="article-detail-page">
      <TwoColumnLayout
        main={<ArticleContent onHeadingsChange={setHeadings} />}
        aside={<ArticleTOC headings={headings} />}
      />
    </div>
  );
};

export default ArticleDetailPage;
