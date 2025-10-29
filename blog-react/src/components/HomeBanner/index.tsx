import { useEffect, useState } from "react";
import { Swiper, SwiperSlide } from "swiper/react";
import { Autoplay, Pagination } from "swiper/modules";
import "swiper/css";
import "swiper/css/bundle";
import "./index.scss";

import { advertisementInfo, type Advertisement } from "@/api/advertisement";
import AdvertisementSlide from "./AdvertisementSlide";

const HomeBanner = () => {
  const [list, setList] = useState<Advertisement[]>([]);

  useEffect(() => {
    advertisementInfo().then((res) => {
      setList(res.data.list);
    });
  }, []);

  if (list.length === 0) return null;

  return (
    <div className="home-banner-wrapper">
      {" "}
      <Swiper
        className="banner-swiper"
        loop
        modules={[Autoplay, Pagination]}
        autoplay={{ delay: 4000 }}
        pagination={{ clickable: true }}
      >
        {" "}
        {list.map((item) => (
          <SwiperSlide key={item.id}>
            {" "}
            <AdvertisementSlide item={item} />{" "}
          </SwiperSlide>
        ))}{" "}
      </Swiper>{" "}
    </div>
  );
};

export default HomeBanner;
