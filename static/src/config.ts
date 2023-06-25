import enUS from "@/locales/en-US.json";
import type { IntlDateTimeFormat } from "vue-i18n";

const datetimeFormats: { [x: string]: IntlDateTimeFormat } = {
  "en-US": {
    month: {
      month: "long",
      year: "numeric",
    },
    day: {
      weekday: "long",
      day: "numeric",
      month: "long",
      year: "numeric",
    },
  },
  "fr-FR": {
    month: {
      month: "long",
      year: "numeric",
    },
    day: {
      weekday: "long",
      day: "numeric",
      month: "long",
      year: "numeric",
    },
  },
};

export const Config = {
  basePath: "/api",
  codeResend: 2 * 60,

  locale: {
    supported: ["en-US", "fr-FR"],
    default: "en-US",
    fallback: <Record<string, string | undefined>>{ fr: "fr-FR", en: "en-US" },
    defaultMsg: { "en-US": enUS },
    datetimeFormats: datetimeFormats,
  },
};
