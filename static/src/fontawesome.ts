import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faBars,
  faGear,
  faXmark,
  faFaceLaughBeam,
  faFaceSmile,
  faFaceMeh,
  faFaceFrown,
  faFaceSadTear,
} from "@fortawesome/free-solid-svg-icons";
import type { App } from "vue";

library.add(
  faGear,
  faBars,
  faXmark,
  faFaceLaughBeam,
  faFaceSmile,
  faFaceMeh,
  faFaceFrown,
  faFaceSadTear
);

export default function setupFontAwesome(app: App) {
  app.component("font-awesome-icon", FontAwesomeIcon);
}
