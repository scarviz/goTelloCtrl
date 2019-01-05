# goTelloCtrl
[DJI Tello](https://store.dji.com/jp/product/tello?vid=45701)をコントロールするためのアプリ。  
[gobot](https://gobot.io)を使用。  

## 環境構築    
Telloのライブラリとして[gobot](https://gobot.io/documentation/platforms/tello/)、
PS4コントローラー等のジョイスティックを扱うため[go-sdl2](https://github.com/veandco/go-sdl2)、Telloのカメラ映像を表示するため[mplayer](http://www.mplayerhq.hu/design7/dload.html)が必要。  

go-sdl2をWindows環境で利用するには、以下を実施する。詳細は[go-sdl2](https://github.com/veandco/go-sdl2)のReadmeにも記載されている。
1. [TDM-GCC](http://tdm-gcc.tdragon.net/download)をインストール
2. [SDL2](http://libsdl.org/download-2.0.php)の「Development Libraries」の「MinGW 32/64-bit」版をダウンロードし、解凍
3. SDL2のx86_64-w64-mingw32フォルダの中身を、TDM-GCCインストール先のx86_64-w64-mingw32フォルダの中に追加
4. TDM-GCCのbinフォルダと、TDM-GCCのx86_64-w64-mingw32フォルダのbinフォルダのパスを環境変数に設定

## PS4コントローラー用設定ファイル  
[gobotのjoystick](https://gobot.io/documentation/platforms/joystick/)を利用しているが、PS4用設定(dualshock4)があっていない。  
そのため「joystick_ps4.json」という設定ファイルを用意して、これを使用している。  

## PS4コントローラー操作  
### △ボタン  
離陸する。  

### 〇ボタン  
着地する。  

### □ボタン  
バックフリップ。  

### ×ボタン  
アプリを終了する。Telloが離陸している場合は着地させる。  

### 左アナログスティック  
縦方向(y方向)で昇降、横方向(x方向)で旋回をする。  

### 右アナログスティック  
縦方向(y方向)で前方後方移動、横方向(ⅹ方向)で左右移動する。  

## WebAPI操作  
http通信でTelloを操作できるようにサーバを立てている。  
使用ポートは8880にしている。  
http://localhost:8880 で以下のパスで、リクエストを出して操作することができる。  

サンプルをres\TelloCtrlに格納している。  
このサンプルを利用する場合は、[bootstrap](https://getbootstrap.com/)の[cssとjs](https://getbootstrap.com/docs/4.2/getting-started/download/#compiled-css-and-js)のフォルダの中身を、res\TelloCtrlのcssとjsフォルダに追加すること。

### /battery  
バッテリー残量を取得する。%表示。  

### /height  
高度を取得する。m表示。  

### /takeoff  
離陸する。  

### /land  
着地する。  

### /palmland  
手のひら着地する。  

### /up  
上昇する。  

### /down  
降下する。  

### /forward  
前方へ移動する。  

### /backward  
後方へ移動する。  

### /left  
左へ移動する。  

### /right  
右へ移動する。  

### /turnright  
右旋回する。  

### /turnleft  
左旋回する。  