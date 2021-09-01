import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_sign_in/google_sign_in.dart';

final authServiceProvider = ChangeNotifierProvider<AuthService>((ref) {
  return AuthService();
});

class AuthService extends ChangeNotifier {
  void showInfo() {
    print('Hello');
  }

  void signIn() async {
    final googleSignIn = GoogleSignIn();
    final firebaseAuth = FirebaseAuth.instance;
    var googleCurrentUser = googleSignIn.currentUser;
    try {
      if (googleCurrentUser == null)
        googleCurrentUser = await googleSignIn.signIn();
      if (googleCurrentUser == null) return null;
      final googleAuth = await googleCurrentUser.authentication;
      final authCredential = GoogleAuthProvider.credential(
          accessToken: googleAuth.accessToken, idToken: googleAuth.idToken);
      final user =
          (await firebaseAuth.signInWithCredential(authCredential)).user;
      if (user != null) {
        print('email: ${user.email}');
      }
    } catch (e) {
      print('error: ${e.toString()}');
    }
  }

  void signOut() async {
    await FirebaseAuth.instance.signOut();
    await GoogleSignIn().signOut();
  }
}
