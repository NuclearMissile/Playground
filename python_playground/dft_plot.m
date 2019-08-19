csvread("pi.csv");

NN = length(xy);
N = 65536;

xx = xy(:,1);
yy = xy(:,2);

figure(1);
plot(xx, yy, '.');

XX = fft(xx)/NN;
YY = fft(yy)/NN;

X = [XX(1+0:1+(NN/2-1)); 0.5*XX(1+NN/2); zeros(N-NN-1, 1); 0.5*XX(1+NN/2); XX(1+(NN/2+1):1+(NN-1))];
Y = [YY(1+0:1+(NN/2-1)); 0.5*YY(1+NN/2); zeros(N-NN-1, 1); 0.5*YY(1+NN/2); YY(1+(NN/2+1):1+(NN-1))];

x = N*ifft(X);
y = N*ifft(Y);

figure(1);
hold on;
plot(x, y, 'b');
hold off;