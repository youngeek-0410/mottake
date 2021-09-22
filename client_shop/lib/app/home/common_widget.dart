import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

Widget errorDialog(String content, VoidCallback onPressed) {
  return AlertDialog(
    title: const Text("Error"),
    content: Text(content),
    actions: [TextButton(onPressed: onPressed, child: const Text("OK"))],
  );
}
