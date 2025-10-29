export const isEmpty = (str?: string) => {
  return str === null || str === undefined || str === "";
};

export const isBlank = (str?: string) => {
  return isEmpty(str?.trim());
};
