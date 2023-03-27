import { Dropdown, MenuProps, Space } from 'antd';
import { useTranslation } from 'react-i18next';

import { localLanguage, lookupLocalStorage } from '../../config/locales';

export const languageMenuItem: MenuProps['items'] = [
  { label: 'English', key: 'en' },
  { label: '简体中文', key: 'zh' }
];

export default function ChangeLanguage({ innerStyle }: { innerStyle: string | undefined }): JSX.Element {
  const { i18n } = useTranslation();

  return (
    <Dropdown
      trigger={['hover']}
      menu={{
        items: languageMenuItem,
        onClick: ({ key }) => {
          i18n.changeLanguage(key);
          localStorage.setItem(lookupLocalStorage, key);
          window.location.reload();
        }
      }}
    >
      <a onClick={(e) => e.preventDefault()}>
        <Space className={innerStyle}>{localLanguage === 'en' ? 'English' : '中文'}</Space>
      </a>
    </Dropdown>
  );
}
