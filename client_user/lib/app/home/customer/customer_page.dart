import 'package:client_user/app/home/models/customer.dart';
import 'package:client_user/app/providers.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:http/http.dart' as http;
import 'package:client_user/constants/urls.dart';
import 'package:client_user/app/home/common_widget.dart';

final customerProvider = FutureProvider<Customer>(
  (ref) => ref.read(databaseProvider)!.getCustomer(),
);

class CustomerDeletePage extends ConsumerStatefulWidget {
  @override
  _CustomerDeletePageState createState() => _CustomerDeletePageState();
}

class _CustomerDeletePageState extends ConsumerState<CustomerDeletePage> {
  @override
  Widget build(BuildContext context) {
    final auth = ref.watch(authServiceProvider);
    final _customerProvider = ref.watch(customerProvider);
    return Scaffold(
        body: Center(
            child: Column(children: [
      _customerProvider.when(data: (data) {
        return ListView(
          shrinkWrap: true,
          children: [_listItem(data.name!, Icon(Icons.person))],
        );
      }, loading: () {
        return const CircularProgressIndicator();
      }, error: (error, stackTrace) {
        return const CircularProgressIndicator();
      }),
      ElevatedButton(
        child: const Text("Sign out"),
        onPressed: auth.signOut,
      ),
      // --------------------------------------------
      // DON'T REMOVE ↓ !!!!!!
      //
      // ElevatedButton(
      //     child: const Text("Delete User"),
      //     onPressed: () {
      //       showDialog(
      //           context: context,
      //           builder: (_) {
      //             return AlertDialog(
      //               title: Text("警告"),
      //               content: Text("本当にユーザを削除しますか？この作業は取り消せません。"),
      //               actions: <Widget>[
      //                 // ボタン領域
      //                 FlatButton(
      //                   child: Text("Cancel"),
      //                   onPressed: () => Navigator.pop(context),
      //                 ),
      //                 FlatButton(
      //                     child: Text("OK"),
      //                     onPressed: () {
      //                       _submit();
      //                     }),
      //               ],
      //             );
      //           });
      //     })
      // ---------------------------------------------
    ])));
  }

  Widget _listItem(String title, Icon icon) {
    return ListTile(title: Text(title), leading: icon);
  }

  Future<void> _submit() async {
    try {
      final database = ref.read(databaseProvider)!;
      await database.deleteCustomer();
      Navigator.of(context).pop();
    } catch (e) {
      showDialog(
          context: context,
          builder: (_) {
            return errorDialog("Failed to delete user.", () {
              Navigator.of(context).pop();
            });
          });
    }
  }
}
