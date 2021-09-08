import 'package:client_user/services/auth.dart';
import 'package:client_user/shared/message_data.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class HomePage extends ConsumerWidget {
  @override
  Widget build(BuildContext context, ScopedReader watch) {
    final authService = watch(authServiceProvider);
    return Scaffold(
      body: watch(messageProvider).when(data: (data) {
        return Scaffold(
            body: Center(
                child: Column(
          children: [
            Text('${data.message}'),
            ElevatedButton(
                onPressed: authService.signOut, child: Text('Sign Out'))
          ],
          mainAxisAlignment: MainAxisAlignment.center,
        )));
      }, loading: () {
        return CircularProgressIndicator();
      }, error: (error, stackTrace) {
        return Text('${error.toString()}');
      }),
    );
  }
}
