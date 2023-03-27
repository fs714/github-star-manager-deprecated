import { BookOutlined, ProfileOutlined, SettingOutlined } from '@ant-design/icons';
import { Layout, Menu, Tooltip } from 'antd';
import { useEffect, useMemo, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { Link, Outlet, useLocation, useNavigate } from 'react-router-dom';

import { GithubBlackImg, GithubWhiteImg, ReactImg } from '../assets';
import ChangeLanguage from '../components/changeLanguage';
import Settings from '../components/setting';
import { useThemeContext } from '../context/theme';

const { Header, Content, Footer } = Layout;

export interface ILayoutStyle {
  headerStyle?: string;
  changeLanguageStyle?: string;
}

export default function AppLayout(): JSX.Element {
  const navigate = useNavigate();
  const location = useLocation();
  const [selectedKey, setSelectedKey] = useState<string[]>([]);
  const [showDrawer, setShowDrawer] = useState<boolean>(false);
  const { theme } = useThemeContext();
  const { menuStyle } = theme;

  const { t } = useTranslation();

  useEffect(() => {
    console.log(theme);
    if (location.pathname === '/') {
      setSelectedKey(['/repositories']);
    } else {
      setSelectedKey([location.pathname]);
    }
  }, [location]);

  const onSelect = (e: any) => {
    if (e?.key) {
      navigate(e.key);
    }
  };

  const layoutStyle: ILayoutStyle = useMemo((): ILayoutStyle => {
    let headerStyle: ILayoutStyle['headerStyle'] = '';
    let changeLanguageStyle: ILayoutStyle['changeLanguageStyle'] = '';

    if (menuStyle === 'light') {
      headerStyle = 'flex h-12 w-screen items-center justify-between px-6 leading-[3rem] bg-white';
      changeLanguageStyle = 'font-bold text-black hover:text-black';
    } else {
      headerStyle = 'flex h-12 w-screen items-center justify-between px-6 leading-[3rem]';
      changeLanguageStyle = 'font-bold text-white/[.65] hover:text-white';
    }

    return { headerStyle, changeLanguageStyle };
  }, [theme]);

  return (
    <div>
      <Layout className="relative flex h-screen w-screen flex-col bg-gray-100">
        <Header className={layoutStyle.headerStyle}>
          <div className="flex justify-start">
            <div className="flex items-center bg-inherit">
              <Link to="/">
                <img src={ReactImg} className="flex" />
              </Link>
            </div>
            <div>
              <Menu
                theme={menuStyle}
                mode="horizontal"
                selectedKeys={selectedKey}
                onSelect={onSelect}
                className="h-12 bg-inherit px-6 font-bold"
                items={[
                  {
                    key: '/repositories',
                    label: t('layout.REPOSITORY'),
                    icon: <ProfileOutlined />
                  },
                  {
                    key: '/tags',
                    label: t('layout.TAGS'),
                    icon: <BookOutlined />
                  }
                ]}
              />
            </div>
          </div>
          <div className="flex justify-start">
            <div>
              <Menu
                theme={menuStyle}
                mode="horizontal"
                selectedKeys={selectedKey}
                onSelect={() => setShowDrawer(true)}
                className="h-12 bg-inherit pr-3 font-bold"
                items={[
                  {
                    key: '/settings',
                    label: t('layout.SETTINGS'),
                    icon: <SettingOutlined />
                  }
                ]}
              />
            </div>
            <div className="bg-inherit pr-6">
              <ChangeLanguage innerStyle={layoutStyle.changeLanguageStyle} />
            </div>
            <div className="flex items-center bg-inherit">
              <Tooltip title="Github">
                <Link to="https://github.com/" target="_blank">
                  <img src={menuStyle === 'light' ? GithubBlackImg : GithubWhiteImg} className="flex h-8 w-8" />
                </Link>
              </Tooltip>
            </div>
          </div>
        </Header>
        <Content className="grow bg-white m-3">
          <Outlet />
        </Content>
        <Footer className="text-center mb-2 p-1 font-bold text-black/[.45]">Copyright © 2023 by fs714</Footer>
      </Layout>
      <Settings setShowDrawer={setShowDrawer} open={showDrawer} onClose={() => setShowDrawer(false)} />
    </div>
  );
}
