import type { InitOptions } from 'i18next';
import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';

import en from './en';
import zh from './zh';

const resources: InitOptions['resources'] = { en: { translation: en }, zh: { translation: zh } };

export type ILanguage = 'en' | 'zh';

export const lookupLocalStorage = 'i18nextLng';
export const localLanguage: ILanguage = (localStorage.getItem(lookupLocalStorage) === 'zh' && 'zh') || 'en';

i18n.use(initReactI18next).init({
  resources,
  lng: localLanguage,
  fallbackLng: 'en',
  interpolation: {
    escapeValue: false
  }
});

export default i18n;
