% Laplacian Growth Model
clc
clear all
close all

w = 50; % width
h = 100; % height
eta = 1;

ni = floor(w/2)*h + (1:3)'; % negative charge
pi = (1:w)'*h; % positive charge

S1 = bsxfun(@plus, ((1:(w-1)) - 1)*h, (1:h)');
D1 = bsxfun(@plus, ((2:w) - 1)*h, (1:h)');
S2 = bsxfun(@plus, ((1:w) - 1)*h, (1:(h-1))');
D2 = bsxfun(@plus, ((1:w) - 1)*h, (2:h)');

S = [S1(:); S2(:)];
D = [D1(:); D2(:)];
A = sparse([S D], [D S], 1);
n = size(A,1);
% Lp = speye(h*w) - bsxfun(@rdivide, A, sum(A,2));
% Lp = spdiags(sum(A,2), 0, n, n) - A;
% rhs = zeros(n, 1);

[X0, Y0] = ind2sub([h,w], (1:h*w)');
gaussian=fspecial('gaussian',[3 3],2);
iter = 1;
frame = 1;

while true
    phi = solve_equation(w, h, ni, pi, 5000);
    [~, adj] = find(A(ni,:));
    adj = unique(adj);
    adj = adj(~ismember(adj, ni));
    k = randsample(adj, 1, true, phi(adj).^eta/sum(phi(adj).^eta));
    ni = [ni; k];

    if mod(iter,10) == 1
        [X, Y] = ind2sub([h,w], ni);
        if max(X) >= h - 1
            break;
        end
        [~, dis] = knnsearch([X,Y], [X0,Y0]);
        im = reshape(1-dis/2, [h,w]);
%         im = zeros(h,w);
%         im(ni) = 1;
%         im = imfilter(im, gaussian);
        imshow(im);
        drawnow;
        Frame(frame) = getframe(gcf);
        frame = frame + 1;
    end
    fprintf('iter: %d\n', iter);
    iter = iter + 1;
end