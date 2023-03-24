import { Layout, Menu } from "antd";
import { ProfileOutlined, BookOutlined, SettingOutlined, GithubOutlined } from "@ant-design/icons"
import { Link, Outlet, useLocation, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { ReactImg, GithubImg } from "../assets";


const { Header, Content, Footer } = Layout

export default function AppLayout(): JSX.Element {
    const navigate = useNavigate()
    const location = useLocation()
    const [selectedKey, setSelectedKey] = useState<string[]>([])

    useEffect(() => {
        if (location.pathname === "/") {
            setSelectedKey(["/repositories"])
        } else {
            setSelectedKey([location.pathname])
        }
    }, [location])

    const onSelect = (e: any) => {
        if (e?.key) {
            navigate(e.key)
        }
    }

    return (
        <Layout className="bg-gray-100">
            <Header className="h-12 leading-[3rem] px-6">
                <div>
                    <div className="float-left pr-6 py-2">
                        <Link to="/">
                            <img src={ReactImg} />
                        </Link>
                    </div>
                    <div className="float-left">
                        <Menu
                            theme="dark"
                            mode="horizontal"
                            selectedKeys={selectedKey}
                            onSelect={onSelect}
                            className="float-left font-bold h-12"
                            items={[
                                {
                                    key: "/repositories",
                                    label: "Repository",
                                    icon: <ProfileOutlined />
                                },
                                {
                                    key: "/tags",
                                    label: "Tags",
                                    icon: <BookOutlined />
                                }
                            ]}
                        />
                    </div>
                    <div className="float-right">
                        <Menu
                            theme="dark"
                            mode="horizontal"
                            selectedKeys={selectedKey}
                            onSelect={onSelect}
                            className="float-left font-bold h-12"
                            items={[
                                {
                                    key: "/settings",
                                    label: "Settings",
                                    icon: <SettingOutlined />
                                }
                            ]}
                        />
                        <div className="float-right pl-6 py-2">
                            <Link to="https://github.com/" target="_blank">
                                <img src={GithubImg} className="h-8 w-8" />
                            </Link>
                        </div>
                    </div>
                </div>
            </Header>
            <Content className="bg-white m-3 min-h-[80vh]">
                <Outlet />
            </Content>
            <Footer className="text-center">
                FS714 Design ©2023 Created by FS714
            </Footer>
        </Layout>
    )
}
