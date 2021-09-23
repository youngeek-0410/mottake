import 'package:client_user/app/home/map/map_page.dart';
import 'package:client_user/app/home/purchase/purchase_page.dart';
import 'package:client_user/app/home/customer/customer_page.dart';
import 'package:client_user/constants/strings.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class TabInfo {
  String label;
  Widget widget;
  TabInfo(this.label, this.widget);
}

class HomePage extends StatelessWidget {
  final List<TabInfo> _tabs = [
    TabInfo("user", CustomerDeletePage()),
    TabInfo("map", MapPage()),
    TabInfo("purchase", PurchasePage()),
  ];

  HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: _tabs.length,
      child: Scaffold(
        appBar: AppBar(
          title: const Text(Strings.appNeme),
          bottom: PreferredSize(
            child: TabBar(
              isScrollable: true,
              tabs: _tabs.map((TabInfo tab) {
                return Tab(text: tab.label);
              }).toList(),
            ),
            preferredSize: const Size.fromHeight(30.0),
          ),
        ),
        body: TabBarView(children: _tabs.map((tab) => tab.widget).toList()),
      ),
    );
  }
}
